<?php
namespace app\api\services;

use app\api\models\{{.tableName}};

class {{.tableName}}Service extends BaseService
{
    public function __construct() 
    {
        $this->model = new {{.tableName}}();
    }

}