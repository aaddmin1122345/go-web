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

    // 登录页面或者需要登录的页面保存当前页面地址
    localStorage.setItem("lastVisitedPage", document.referrer);

// 登录成功后从 localStorage 获取保存的页面地址并重定向
    $.ajax({
        type: "POST",
        url: "/api/login",
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        data: JSON.stringify(jsonData),
        success: function (data) {
            if (data.valid) {
                alert("登录成功");
                const redirectTo = localStorage.getItem("lastVisitedPage") || "/";
                window.location.href = redirectTo;
            }
        },
        error: function (err) {
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
