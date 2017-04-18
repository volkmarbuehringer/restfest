package generator

import "os"

const BindVar string = "$"
const EqBindVar string = "=$"
const dbTimestamp string = "current_timestamp"

var dbschema = os.Getenv("PGSCHEMA")

func dbSequenzer(tab string) string {
	return ""
}

const transform_sql string = `
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

var sqlfunctionparams string = `sELECT parameter_name,'*'||` + transform_sql + ` as coltrans
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
select flag,table_name,column_name,routine_name,
(select case when count(*) > 0 and flag <> 3 then true else false end from flagger where name = x.table_name) as tflag
from(
select
case when  t.table_name in ( 'weburl') then 4 else 1 end as flag,
t.table_name,pk.column_name,
'' as routine_name,t.table_schema
from
information_schema.tables t
inner join pk on ( t.table_name = pk.table_name)
 where t.table_schema ='` + dbschema + `'
   union all
select 2,c.table_name
,column_name
,'' as routine, v.table_schema
from information_schema.views v
inner join information_schema.columns c on v.table_name = c.table_name and ordinal_position = 1
where v.table_schema ='` + dbschema + `' and c.table_schema = '` + dbschema + `'
union all
sELECT 3,routines.type_udt_name,routine_name,specific_name, specific_schema as table_schema
 FROM information_schema.routines
    WHERE routines.specific_schema='` + dbschema + `'
		and data_type = 'USER-DEFINED'
		and routine_type ='FUNCTION'
	) x
	where table_name not like 'hst%' and table_name not like '%hst'
	order by table_name
	limit 400
`

var sqlallcols string = `select column_name,case when is_nullable = 'YES' and data_type not in ('USER-DEFINED','ARRAY')then
'*' else '' end||` + transform_sql + `as coltrans,
column_name
from information_schema.columns
where table_name =$1 and table_schema = '` + dbschema + `' and column_name not like '$%' order by  ordinal_position `
