<?php
namespace app\{{.module}}\services;

use app\{{.module}}\models\{{.modelName}};

class {{.fileName}}Service extends BaseService
{
    public function __construct() 
    {
        $this->model = new {{.modelName}}();
    }

}