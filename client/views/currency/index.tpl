<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<link href="/static/css/table-responsive.css" rel="stylesheet">
</head><body class="sticky-header">
<section> {{template "inc/left.tpl" .}}
  <!-- main content start-->
  <div class="main-content" >
    <!-- header section start-->
    <div class="header-section">
      <!--toggle button start-->
      <a class="toggle-btn"><i class="fa fa-bars"></i></a>
      <!--toggle button end-->
      <!--search start-->      
      <!--search end-->
      {{template "inc/user-info.tpl" .}} 
    </div>
    <!-- header section end-->
    <!-- page heading start-->
    <!-- page heading end-->
    <!--body wrapper start-->
    <div class="wrapper">
      <div class="row">
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 我的币&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="/currency/create">+ 创建</a>
              <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              </span>
            </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="funds-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>币名</th>
                        <th>创建日期</th>
                        <th>总数</th>
                        <th>剩余</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    {{range $k,$v := .myCurrencys}}
                    <tr>
                      <td>{{$v.ID}}{{if eq $v.ID "CNY" "USD"}}(默认){{end}}</td>
                      <td>{{$v.CreateDate}}</td>
                      <td>{{$v.Count}}</td>
                      <td>{{$v.LeftCount}}</td>
                      <td>{{if eq $v.ID "CNY" "USD"}}{{else}}<a href="/currency/release/{{$v.ID}}">增加</a>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{{if gt $v.LeftCount 0.0}}<a href="/currency/assign/{{$v.ID}}">分发</a>{{end}}{{end}}</td>                     
                    </tr>
                    {{end}}
                    </tbody>
                  </table>
                </form>
				      </section>
            </div>
          </section>
        </div>
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 我的资产
              <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              </span>
            </header>
            <div class="panel-body">
              <section id="unseen">
                <form id="funds-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>币名</th>
                        <th>数量</th>
                        <th>锁定量</th>
                      </tr>
                    </thead>
                    <tbody>
                    {{range $k,$v := .myAssets}}
                    <tr>
                      <td>{{$v.Currency}}</td>
                      <td>{{$v.Count}}</td>
                      <td>{{$v.LockCount}}</td>                    
                    </tr>
                    {{end}}
                    </tbody>
                  </table>
                </form>
				      </section>
            </div>
          </section>
        </div>
        <div class="col-sm-12">
          <section class="panel">
            <header class="panel-heading"> 挂单
              <span class="tools pull-right"><a href="javascript:;" class="fa fa-chevron-down"></a>
              </span>
            </header>
            <div class="panel-body">
              <section id="unseen">
                <form class="form-horizontal adminex-form" id="exchange-form">
                <div class="form-group">
                  <div class="col-sm-1" style="width:10%">
                    <select name="srcCurrency" class="form-control">
                    {{range $k,$v := .myCurrencyIds}}
                    <option value={{$v}}>{{$v}}</option>
                    {{end}}
                    </select>
                  </div>
                  <div class="col-sm-1" style="width:10%">
                    <input type="number" name="srcCount" class="form-control" placeholder="数量">
                  </div>
                  <label class="col-sm-1 control-label" style="text-align:center">兑换</label>
                  <div class="col-sm-1" style="width:10%">
                    <select name="desCurrency" class="form-control">
                     {{range $k,$v := .allCurrencyIds}}
                    <option value={{$v}}>{{$v}}</option>
                    {{end}}
                    </select>
                  </div>
                  <div class="col-sm-1" style="width:10%">
                    <input type="number" name="desCount" class="form-control" placeholder="数量">
                  </div>
                  <label class="col-sm-1 control-label" style="text-align:center"><input name="isBuyAll" type="radio" value="true" checked/>&nbsp;&nbsp;买完 </label> 
                  <label class="col-sm-1 control-label" style="text-align:center"><input name="isBuyAll" type="radio" value="false" />&nbsp;&nbsp;卖完 </label> 
                  <div class="col-sm-1">
                    <button type="submit" class="btn btn-primary">挂单</button>
                  </div>
                </div>
                <div class="panel-body">
              <section id="unseen">
                <form id="tx-form-list">
                  <table class="table table-bordered table-striped table-condensed">
                    <thead>
                      <tr>
                        <th>单号</th>
                        <th>原单</th>
                        <th>源币</th>
                        <th>目标币</th>
                        <th>类型</th>
                        <th>状态</th>
                        <th>挂单时间</th>
                        <th>成交时间</th>
                        <th>实际成交</th>
                        <th>操作</th>
                      </tr>
                    </thead>
                    <tbody>
                    {{range $k,$v := .txs}}
                    <tr>
                      <td>{{$v.UUID}}</td>
                      <td>{{if ne $v.RawUUID $v.UUID}}{{$v.RawUUID}}{{end}}</td>
                      <td>{{$v.SrcCurrency}}/{{$v.SrcCount}}</td>
                      <td>{{$v.DesCurrency}}/{{$v.DesCount}}</td>
                      <td>{{if $v.IsBuyAll}}买完{{else}}卖完{{end}}</td>
                      <td>{{if eq $v.Status 0}}待交易{{else if eq $v.Status 1}}完成{{else if eq $v.Status 2}}过期{{else if eq $v.Status 3}}撤单{{else}}{{end}}</td>
                      <td>{{$v.PendedDate}}</td>
                      <td>{{$v.FinishedDate}}</td>
                      <td>{{$v.FinalCost}}/{{$v.DesCount}}</td>
                      <td>{{if eq $v.Status 0}}<a data-id="{{$v.UUID}}" class="cancel-class">撤单</a>{{end}}</td>
                    </tr>
                    {{end}}
                    </tbody>
                  </table>
                </form>
				      </section>
            </div>
          </section>
        </div>
      </div>
    </div>
    <!--body wrapper end-->
    <!--footer section start-->
    {{template "inc/foot-info.tpl" .}}
    <!--footer section end-->
  </div>
  <!-- main content end-->
</section>
{{template "inc/foot.tpl" .}}
</body>
</html>
