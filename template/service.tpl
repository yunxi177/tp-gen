<?php
namespace app\{{.cfg.Module}}\{{.cfg.ServiceName}}{{- if ne .origin  ""}}\{{.origin}}{{- end}};
{{- if ne .origin ""}}
use app\api\{{.cfg.ServiceName}}\BaseService;
{{- end}}
use app\{{.cfg.Module}}\{{.cfg.ModelName}}\{{.modelName}};

class {{.fileName}}Service extends BaseService
{
    public function __construct() 
    {
        $this->model = new {{.modelName}}();
    }

}