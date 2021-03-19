<!DOCTYPE HTML>
<!--
	Paradigm Shift by HTML5 UP
	html5up.net | @ajlkn
	Free for personal and commercial use under the CCA 3.0 license (html5up.net/license)
-->
<html>
	<head>
        {% block head %}
		<title>{{ Title }}</title>
		<meta charset="utf-8" />
		<meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no" />
		<meta name="description" content="" />
		<meta name="keywords" content="" />
		<link rel="stylesheet" href="/public/css/main.css" />
                {% endblock %}
	</head>
	<body class="is-preload">
        {% block body %}
{% endblock %}
		<!-- Scripts -->
			<script src="/public/js/jquery.min.js"></script>
			<script src="/public/js/jquery.scrolly.min.js"></script>
			<script src="/public/js/browser.min.js"></script>
			<script src="/public/js/breakpoints.min.js"></script>
			<script src="/public/js/util.js"></script>
			<script src="/public/js/main.js"></script>

	</body>
</html>
