package generator

import "os"

const BindVar string = "$"
const EqBindVar string = "=$"
const dbTimestamp string = "current_timestamp"

var dbschema = os.Getenv("PGSCHEMA")

func dbSequenzer(tab string) string {
	return ""
}

const transformSQL string = `
case
when data_type ='ARRAY' then
case when ltrim(udt_name,'_') not in
('int2','int4','int8') then 'Ar' else  '[]' end
 else '' end||
case  coalesce( case
when data_type in ('USER-DEFINED','ARRAY') then ltrim(udt_name,'_') end,data_type)
when   'int2' then
'int16'
when   'int4' then
'int32'
when   'int8' then
'int64'
when 'integer' then
'int32'
when 'bigint' then
'int64'
when 'smallint' then
'int16'
when  'boolean' then
'bool'
when 'double precision' then
'float64'
when 'real' then
'float64'
when 'json' then
'map[string]interface{}'
when 'jsonb' then
'map[string]interface{}'
when 'character varying' then 'string'
when 'text' then 'string'
when 'character' then
'string'
when 'json' then
'string'
when  'numeric'  then
'string'
when 'timestamp without time zone'
then 'time.Time'
when 'timestamp with time zone' then 'time.Time'
when 'date' then 'time.Time'
when  'uuid' then 'string'
when  'jsonb' then 'string'
else  upper(substr(coalesce( case when data_type in ('USER-DEFINED','ARRAY') then ltrim(udt_name,'_') end,data_type),1,1))||
substr(coalesce( case when data_type in ('USER-DEFINED','ARRAY') then ltrim(udt_name,'_') end,data_type),2)
 end `

var sqlfunctionparams string = `sELECT parameter_name,'*'||` + transformSQL + ` as coltrans
FROM  information_schema.parameters where parameters.specific_name =$1
and parameters.specific_schema='` + dbschema + `'
and parameter_mode = 'IN' and parameter_name is not null
ORDER BY  parameters.ordinal_position`

var sqlalltabs string = `
with pk as (
	SELECT
	  tc.table_name
		,ccu.column_name
	 FROM
	 information_schema.table_constraints
	 tc JOIN
	 information_schema.constraint_column_usage
	AS
	 ccu USING
	(constraint_schema,
	 constraint_name)
		where
	 constraint_type =
	'PRIMARY KEY'
	and tc.table_schema ='` + dbschema + `'
), flagger as (
	select distinct substr(udt_name,2) as name from information_schema.columns c  where data_type in ('USER-DEFINED','ARRAY')
	and c.table_schema ='` + dbschema + `'
	and substr(udt_name,2) in (
		select  t.typname
		from pg_type t
		where (
			  t.typtype   not in('b', 'p', 'r')
		    and t.typarray > 0
			)
		    and t.typnamespace = ( select oid from pg_namespace where nspname = '` + dbschema + `')
	)
)
select flag,x.table_name,column_name,case when rower = 1 then routine_name else specific_name end,
(select case when count(*) > 0 and flag <> 3 then true else false end from flagger where name = x.table_name) as tflag,
specific_name
from(
select
case when  t.table_name in ( 'weburl') then 4 else 1 end as flag,
t.table_name,pk.column_name,
'' as routine_name,t.table_schema,t.table_name as specific_name, row_number()over ( partition by t.table_name) as rower
from
information_schema.tables t
inner join pk on ( t.table_name = pk.table_name)
 where t.table_schema ='` + dbschema + `'
 union all
 select 5,t.user_defined_type_name,attribute_name  , '' as routine,t.user_defined_type_schema,t.user_defined_type_name as specific_name,
 row_number()over ( partition by user_defined_type_name) as rower
 from information_schema.user_defined_types t
inner join information_schema.attributes on t.user_defined_type_name=udt_name   and ordinal_position = 1
 where t.user_defined_type_schema ='` + dbschema + `' and udt_schema = '` + dbschema + `'
   union all
select 2,c.table_name
,column_name
,'' as routine, v.table_schema,c.table_name as specific_name,row_number()over(partition by c.table_name) as rower
from information_schema.views v
inner join information_schema.columns c on v.table_name = c.table_name and ordinal_position = 1
where v.table_schema ='` + dbschema + `' and c.table_schema = '` + dbschema + `'
union all
sELECT 3,routines.type_udt_name,routine_name,routine_name, specific_schema as table_schema,specific_name,row_number()over(partition by routine_name) as rower
 FROM information_schema.routines
    WHERE routines.specific_schema='` + dbschema + `'
		and data_type = 'USER-DEFINED'
		and routine_type ='FUNCTION'
	) x
	inner join ` + dbschema + `.testselector ts on project = '` + os.Args[1] + `' and x.table_name like coalesce(ts.table_name,'%')
	where x.table_name not like 'hst%' and x.table_name not like '%hst'
	and x.table_name not like 'pg_stat%'
	order by x.table_name
	limit 400
`

var sqlallcols string = `
select
x.column_name,
case when x.is_nullable = 'YES' and x.data_type not in ('USER-DEFINED','ARRAY')then '*' else '' end||` + transformSQL + `as coltrans,
x.column_name
from
(
select c.column_name,c.ordinal_position, c.table_name,c.is_nullable,c.data_type, c.udt_name
from information_schema.columns c
where c.table_schema = '` + dbschema + `' and c.column_name not like '$%'
union all
select a.attribute_name,a.ordinal_position,a.udt_name as table_name,a.is_nullable,a.data_type,attribute_udt_name as udt_name
from information_schema.attributes a
 where  a.udt_schema = '` + dbschema + `'
 ) x
 where x.table_name =$1
order by  x.ordinal_position`
