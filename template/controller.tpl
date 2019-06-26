<?php
namespace app\api\controllers;

use app\api\services\{{.tableName}}Service as Service;
use app\api\validates\{{.tableName}}Validate as Validate;

class {{.tableName}} extends BaseController
{
    {{if .genCondition.GenAdd }}
    /**
     * 添加数据
     *
     * @return void
     */
    public function add() 
    {
        $data = input('post.');
        $validate = new Validate();
        $validate->scene('add');
        if(! $validate->check($data)) {
            return $this->wrongReturn($validate->getError());
        }
        
        (new Service)->add($data);

        return $this->okReturn([]);
    }
    {{end}}
    {{- if .genCondition.GenDel }}
    /**
     * 删除数据
     *
     * @return void
     */
    public function del() 
    {
        $id = input('get.id');
        
        (new Service)->del([['id', '=', $id]]);
        return $this->okReturn([]);
    }{{- end}}
    {{- if .genCondition.GenUp }}
    /**
     * 更新数据
     *
     * @return void
     */
    public function update() 
    {
        $id = input('get.id');
        $data = input('put.');
    
        (new Service)->update([['id', '=', $id]], $data);
        return $this->okReturn([]);
    }{{- end}}
    {{- if .genCondition.GenList }}
    /**
     * 获取列表
     *
     * @return void
     */
    public function list() 
    {
        $this->getPage();
        $service = new Service;
        $list = $service->all([], ['*'], $this->page, $this->limit);

        return $this->okReturn($list, $service->getPageTotal());
    }{{- end}}

    {{- if .genCondition.GenInfo }}
    /**
     * 获取信息
     *
     * @return void
     */
    public function info() 
    {
        $id = input('get.id');
        if(empty($id)) {
            return $this->wrongReturn('参数错误');
        }

        $info = (new Service)->one([['id', '=', $id]], ['*']);
        $info = empty($info)?[]:$info->toArray();
        return $this->okReturn($info);
    }
    {{- end}}
}