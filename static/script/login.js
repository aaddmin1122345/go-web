function login() {
    const name = $("#loginUserName").val();
    const pwd = $("#loginPassWord").val();
    const rememberMe = $("#rememberMe").is(":checked"); // 修正获取复选框状态的方法

    const jsonData = {
        userName: name,
        passWord: pwd,
        rememberMe: rememberMe // 将复选框的状态添加到发送的数据中
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
                window.location.href = "/";
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

$(document).ready(function () {
    $("#loginForm").submit(function (event) {
        event.preventDefault(); // 阻止表单默认提交行为
        login();
    });
});
