function login() {
    const name = $("#loginUserName").val();
    const pwd = $("#loginPassWord").val();

    const jsonData = {
        userName: name,
        passWord: pwd
    };

    if (name.length === 0 || pwd.length === 0) {
        $("#errorInfo").html("用户名和密码都不能为空").css("color", "red");
        return;
    }

    $.ajax({
        type: "POST",
        url: "/api/login",
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        data: JSON.stringify(jsonData),
        success: function (data) {
            if (data.valid) {
                // 登录成功时将错误信息设置为绿色
                alert("登录成功");
                $("#errorInfo").html("登陆成功").css("color", "green");
                console.log("成功");
            }
        },
        error: function (err) {
            // 错误处理，提示更详细的错误信息
            let errorMessage = "账号或者密码错误";
            if (err.responseJSON && err.responseJSON.message) {
                errorMessage = err.responseJSON.message;
            }
            $("#login_error").html(errorMessage).css("color", "red");
        }
    });
}

$(document).ready(function() {
    $("#loginForm").submit(function(event) {
        event.preventDefault(); // 阻止表单默认提交行为
        login();
    });
});