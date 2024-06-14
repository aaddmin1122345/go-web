$(document).ready(function () {
    $('#registerForm').submit(function (event) {
        event.preventDefault(); // 防止默认表单提交行为
        register();
    });
});

function register() {
    // 获取选中的性别
    const selectedSex = $('input[name="sex"]:checked').val();
    // 获取选中的用户类型
    const selectedUserType = $('input[name="userType"]:checked').val();

    const user = {
        PhoneNum: $('#phoneNum').val(),
        Username: $('#username').val(),
        Sex: selectedSex,
        Email: $('#email').val(),
        Password: $('#password').val(),
        UserType: selectedUserType,
    };

    // 检查密码长度是否大于或等于6位
    if (!/^\S{6,}$/.test(user.Password)) {
        alert("密码必须大于或等于6位");
        return;
    }

    // 如果有字段未填写，可以在这里添加逻辑进行检查
    if (!user.PhoneNum || !user.Username || !user.Sex || !user.Email || !user.Password || !user.UserType) {
        alert("所有字段都必须填写");
        return;
    }

    $.ajax({
        url: '/api/register',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(user),
        success: function () {
            alert('注册成功');
            window.location.href = "/"
        },
        error: function (xhr, status, error) {
            console.log(error);
        }
    });
}
