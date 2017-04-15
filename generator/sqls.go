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
when data_type ='ARRAY' then '[]' else '' end||
case  coalesce( case
when data_type in ('USER-DEFINED','ARRAY') then ltrim(udt_name,'_') end,data_type)
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
when  'numeric'  then
'string'
when 'timestamp without time zone'
then 'time.Time'
when 'timestamp with time zone' then 'time.Time'
when 'date' then 'time.Time'
when  'uuid' then 'string'
when  'jsonb' then 'string'
else  initcap(coalesce( case when data_type in ('USER-DEFINED','ARRAY') then ltrim(udt_name,'_') end,data_type))
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
)
select
1 as flag, t.table_name,pk.column_name,
'' as routine_name
from
information_schema.tables t
inner join pk on ( t.table_name = pk.table_name)
 where t.table_schema ='` + dbschema + `'
   and t.table_name not in ( 'logger')
union all
select 2,c.table_name
,column_name
,'' as routine
from information_schema.views v
inner join information_schema.columns c on v.table_name = c.table_name and ordinal_position = 1
where v.table_schema ='` + dbschema + `' and c.table_schema = '` + dbschema + `'
union all
sELECT 3,routines.type_udt_name,routine_name,specific_name
 FROM information_schema.routines
    WHERE routines.specific_schema='` + dbschema + `'
		and data_type = 'USER-DEFINED'
		and routine_type ='FUNCTION'
`

var sqlallcols string = `select column_name,case when is_nullable = 'YES' then
'*' else '' end||` + transform_sql + `as coltrans,
column_name
from information_schema.columns
where table_name =$1 and table_schema = '` + dbschema + `' and column_name not like '$%' order by  ordinal_position `
