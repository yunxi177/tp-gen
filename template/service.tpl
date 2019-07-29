<?php
namespace app\{{.module}}\services\{{- if ne .origin  ""}}{{.origin}}{{- end}};
{{- if ne .origin ""}}
use app\api\services\BaseService;
{{- end}}
use app\{{.module}}\models\{{.modelName}};

class {{.fileName}}Service extends BaseService
{
    public function __construct() 
    {
        $this->model = new {{.modelName}}();
    }

}