package generator

const BindVar string = "$"
const EqBindVar string = "=$"
const dbTimestamp string = "current_timestamp"
const dbschema = "public"

func dbSequenzer(tab string) string {
	return ""
}

const sqlalltabs string = `SELECT
  tc.table_name
	,c.column_name
 FROM
 information_schema.table_constraints
 tc JOIN
 information_schema.constraint_column_usage
AS
 ccu USING
(constraint_schema,
 constraint_name)
JOIN
 information_schema.columns
AS
 c ON
 c.table_schema
= tc.constraint_schema
AND tc.table_name
= c.table_name
AND ccu.column_name
= c.column_name
where
 constraint_type =
'PRIMARY KEY'   and tc.table_name not in ( 'logger')
 and tc.table_schema ='` + dbschema + `'
union all
select c.table_name
,column_name
from information_schema.views v
inner join information_schema.columns c on v.table_name = c.table_name and ordinal_position = 1
where v.table_schema ='` + dbschema + `' and c.table_schema = '` + dbschema + `'
`

const sqlallcols string = `select column_name,case when data_type = 'numeric' then
case when numeric_scale > 0 then column_name||'::varchar' else
column_name||'::bigint' end else column_name end as column_type,
case data_type
when 'integer' then
case when is_nullable = 'YES' then
'JSONNullInt64'
else
'int64'
end
when 'bigint' then
case when is_nullable = 'YES' then
'JSONNullInt64'
else
'int64'
end
when 'double precision' then
case when is_nullable = 'YES' then
'JSONNullFloat64'
else
'float64'
end
when 'character varying' then
case when is_nullable = 'YES' then
'JSONNullString'
else
'string'
end
when 'character' then
case when is_nullable = 'YES' then
'JSONNullString'
else
'string'
end
when 'text' then
case when is_nullable = 'YES' then
'JSONNullString'
else
'string'
end

when 'numeric'  then
case
  when numeric_scale > 0 and is_nullable = 'YES' then 'JSONNullFloat64'
  when numeric_scale > 0 and is_nullable = 'NO' then 'float64'
  when numeric_scale = 0 and is_nullable = 'NO' then 'int64'
else
'JSONNullInt64'
  end
when 'timestamp without time zone' then
case when is_nullable = 'YES' then 'JSONNullString'
else
  'time.Time'
  end
when 'date' then
  case when is_nullable = 'YES' then 'JSONNullString'
  else
    'time.Time'
    end
	else 'gaga'
end as coltrans
from information_schema.columns
where table_name =$1 and table_schema = '` + dbschema + `' and column_name not like '$%' order by  ordinal_position `
