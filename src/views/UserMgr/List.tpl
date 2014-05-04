

<style type="text/css">
._list_query_input {
    padding: 5px 5px 5px 30px;
    background: url(/ids/~/ids/img/search-16.png) no-repeat 8px; 
    width: 200px;
}
</style>

<div class="ids-user-panel">

<table width="100%">
  <tr>
    <td>
      <form id="eeobsd" action="#" class="form-inlines">
        <input id="query_text" type="text"
          class="form-control _list_query_input" 
          placeholder="Enter to search" 
          value="{{.query_text}}">
      </form>
    </td>
    <td align="right">
      <button type="button" 
        class="btn btn-primary btn-sm" 
        onclick="idsWorkLoader('user-mgr/edit')">
        New User
      </button>
    </td>
  </tr>
</table>
{{if .list}}
<table class="table table-hover">
  <thead>
    <tr>
      <th>Username</th>
      <th>Nickname</th>
      <th>Email</th>
      <th>Timezone</th>
      <th>Status</th>
      <th>Role</th>
      <th>Registered</th>
      <th>Updated</th>
      <th></th>
    </tr>
  </thead>
  <tbody>
    {{range .list}}
    <tr>
      <td>{{.uname}}</td>
      <td>{{.name}}</td>
      <td>{{.email}}</td>
      <td>{{.timezone}}</td>
      <td>{{.status}}</td>
      <td>
        {{if .roles_display}}
        {{range .roles_display}}
        <div>{{.}}</div>
        {{end}}
        {{end}}
      </td>
      <td>{{date .created}}</td>
      <td>{{date .updated}}</td>
      <td>
        <a class="jdiskq" href="#{{.uid}}">Edit</a>
      </td>
    </tr>
    {{end}}
  </tbody>
</table>
{{else}}
<div class="alert alert-info" style="margin:20px 0;">Data not found</div>
{{end}}
</div>


<script type="text/javascript">

$(".jdiskq").click(function() {
    var uid = $(this).attr("href").substr(1);
    idsWorkLoader("user-mgr/edit?uid="+ uid);
});

function _usermgr_list_refresh()
{
    var uri = "query_text="+ $("#query_text").val();

    $.ajax({
        type    : "POST",
        url     : "/ids/user-mgr/list",
        data    : uri,
        timeout : 3000,
        success : function(rsp) {
            $("#work-content").html(rsp);
        },
        error   : function(xhr, textStatus, error) {
            //lessAlert("#azt02e", 'alert-danger', textStatus+' '+xhr.responseText);
        }
    });
}

$("#eeobsd").submit(function(event) {
    event.preventDefault();
    _usermgr_list_refresh();
});

</script>
