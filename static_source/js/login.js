
;(function($){

  'use strict';

  $(function(){
    $('.login-btn').bind('click', function(event) {
      var isVerify = true;

      $('.error-tip').css('opacity', 0);

      $('.form-group').find('input').each(function() {
        if($.trim($(this).val()).length === 0) {
          $(this).parent().find('.error-tip').css('opacity', 1).find('span').html('内容不能为空');
          isVerify = false;
          return false;
        }
      });

      if(isVerify === false) {
        return false;
      }

      $.get('/interface/user/login',  $('.login_form').serialize(), function(data, state) {

        if(data.statuscode === 20010) {
          $('.captcha-error').css('opacity', 1).find('span').html('验证码错误');
        }else if(data.statuscode !== 10000) {
          $('.password-error').css('opacity', 1).find('span').html('帐号或密码错误');
        }else {
          window.location.href = '/';
        }
      });
      return false;
    });
  });

})(jQuery);