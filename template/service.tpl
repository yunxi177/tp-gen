<?php
namespace app\{{.module}}\services;

use app\{{.module}}\models\{{.tableName}};

class {{.tableName}}Service extends BaseService
{
    public function __construct() 
    {
        $this->model = new {{.tableName}}();
    }

}