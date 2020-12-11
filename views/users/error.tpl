<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>{{config "String" "globaltitle" ""}}</title>
{{template "inc/meta.tpl" .}}
<style>
.form-signin .help-block{color:#a94442;}
</style>
</head><body class="login-body">
<div class="container">
  <form class="form-signin" id="login-form">
    <div class="form-signin-heading text-center">
      <h1 class="sign-title">some error!</h1>
    <div class="login-wrap">
      <p>{{.error}}</p>
    </div>
  </form>
</div>
{{template "inc/foot.tpl" .}}
</body>
</html>
