package generator

import "os"

const BindVar string = "$"
const EqBindVar string = "=$"
const dbTimestamp string = "current_timestamp"

var dbschema = os.Getenv("PGSCHEMA")

func dbSequenzer(tab string) string {
	return ""
}

var sqlfunctionparams string = `sELECT parameter_name,
case when data_type
in ('integer') then
'*int32'
when data_type in ('bigint') then
'*int64'
when data_type in ('smallint') then
'*int16'
when data_type in ( 'boolean') then
'*bool'
when data_type in ('double precision','real')  then
'*float64'
when data_type in ('character varying',
'text',
'character') then
--'db.JSONNullString'
'*string'
--'db.JSONString'
when data_type = 'numeric'  then
 '*string'
when data_type in ('timestamp without time zone',
	'timestamp with time zone',
'date') then
'*time.Time'
	else 'gaga'
end as coltrans
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

var sqlallcols string = `select column_name,
case
when data_type ='ARRAY' and udt_name = '_int4' then
'[]int32'
when data_type ='ARRAY' and udt_name = '_int8' then
'[]int64'
when data_type ='ARRAY' and udt_name = '_character varying' then
'[]string'
when data_type
in ('integer') then
case when is_nullable = 'YES' then
'*int32'
else
'int32'
end
when data_type in ('bigint') then
case when is_nullable = 'YES' then
'*int64'
else
'int64'
end
when data_type in ('smallint') then
case when is_nullable = 'YES' then
'*int16'
else
'int16'
end
when data_type in ( 'boolean') then
case when is_nullable = 'YES' then
'*bool'
else
'bool'
end
when data_type in ('double precision','real')  then
case when is_nullable = 'YES' then
'*float64'
else
'float64'
end
when data_type in ('character varying',
'text',
'character') then
case when is_nullable = 'YES' then
'*string'
else
'string'
end
when data_type = 'numeric'  then
case when is_nullable = 'YES' then
'*string'
else
'string'
end
when data_type in ('timestamp without time zone',
	'timestamp with time zone',
'date') then
case when is_nullable = 'YES' then
	'*time.Time'
else
  	'time.Time'
  end
when data_type = 'uuid' then '*string'
when data_type = 'jsonb' then '*string'
	else 'string'
end as coltrans,
/*
case
when data_type in ('character varying',
'text',
'character') then
case when is_nullable = 'YES' then
'coalesce('||column_name||','''')as '||column_name
else
column_name
end
else
column_name
end
*/
column_name
from information_schema.columns
where table_name =$1 and table_schema = '` + dbschema + `' and column_name not like '$%' order by  ordinal_position `
