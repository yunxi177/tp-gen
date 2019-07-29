<?php
namespace app\api\models;

class {{.fileName}} extends {{.origin}}BaseModel
{
    protected $table="{{- if eq .origin "DT"}}pb_{{- end}}{{.tableName}}";
}