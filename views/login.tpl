<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Login Page</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<link rel="stylesheet" href="http://apps.bdimg.com/libs/fontawesome/4.4.0/css/font-awesome.min.css">
	<link rel="stylesheet" href="style/loginregister.css">
</head>
<body>
	<header></header>
	<div class="login-wrap">
		<div class="login-part">
			<h3>账号登陆</h3>
			<div class="login-info">
				<form action="">
					<div class="form-group">
						<input type="text" class="" placeholder="用户名/邮箱/手机号" required>
						<i class="fa fa-user"></i>
					</div>
					<div class="form-group">
						<input type="password" class="" placeholder="请输入密码" required>
						<i class="fa fa-lock"></i>
					</div>
					<div class="form-group">
						<input type="text" class="captcha-input" placeholder="请输入验证码" required>
						<i class="fa fa-lock"></i>
						<img src="http://www.chinwhiz.cn/wechat-subscribe-master/pc/login/php/vdimgck.php" class="captcha-img">
					</div>
					<div class="login-forget-password">
						<span>
							<input type="checkbox" id="rememberMe">
							<label for="rememberMe">下次自动登录</label>
						</span>
						<a href="javascript:;">忘记密码</a>
					</div>
					<input class="login-btn" value="登 录" type="submit">
				</form>
				
			</div>
			<div class="split-line"></div>
			<div class="login-third-part">
				<span>第三方帐号登录</span>
				<a href="javascript:;" class="sina" target="_blank"></a>
				<a href="javascript:;" class="baidu" target="_blank"></a>
				<a href="javascript:;" class="qq" target="_blank"></a>
				<a href="javascript::" class="github" target="_blank"></a>

				<div class="login-register"><span>还没有帐号？<a href="javascript:;" target="_blank">立即注册</a></span>
                
           	</div>
			</div>
		</div>
	</div>
</body>
</html>