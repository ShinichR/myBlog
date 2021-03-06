{{define "header"}}
<html>
  <head>
    <title>ShinichR Blog</title>
    <meta name="keywords" content="{{.keywords}}"/>
    <meta name="description" content="{{.description}}"/>
    <link href='http://fonts.googleapis.com/css?family=Open+Sans:400,700' rel='stylesheet' type='text/css' />
    <link href='/static/blog/css/style.css' rel='stylesheet' type='text/css' />
    <link href='/static/blog/css/syntax.css' rel='stylesheet' type='text/css' />
    <link href='/static/blog/css/responsive.css' rel='stylesheet' type='text/css' />
    <!-- - -->
    <script src='/static/blog/js/jquery.js' type='text/javascript'></script>
    <script src='/static/blog/js/pd.js' type='text/javascript'></script>
    <script src='/static/blog/js/basics.js' type='text/javascript'></script>
    <!-- - -->
    <meta content='width=device-width, initial-scale=1.0, user-scalable=no' name='viewport'>
    <meta content='text/html; charset=utf-8' http-equiv='content-type' />

  </head>
  <body>
    <header>
    <a id="go-back-home" href="/"><img src="/static/blog/img/scribble.png" alt="scribble" width="53" height="59"></a>
   
    </header>

    <div id='container'>
      <div class="block">

        <a target="_top" class="main" href="/">Home</a>

      

      </div>

{{end}}