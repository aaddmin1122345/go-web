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
                alert("登录成功!");
                // 可以在此处添加进一步的成功处理逻辑，例如页面跳转
            } else {
                $("#errorInfo").html("用户名或密码错误").css("color", "red");
            }
        },
        error: function (err) {
            // 错误处理，提示更详细的错误信息
            let errorMessage = "登录失败，请稍后重试。";
            if (err.responseJSON && err.responseJSON.message) {
                errorMessage = err.responseJSON.message;
            }
            $("#errorInfo").html(errorMessage).css("color", "red");
        }
    });
}


function updateUser(){
    $(document).ready(function () {
        // 绑定表单提交事件
        $('#updateUserForm').submit(function (event) {
            event.preventDefault();

            const user = {
                StudID: $('#updateStudID').val(),
                Username: $('#updateUsername').val(),
                Sex: $('#updateSex').val(),
                Email: $('#updateEmail').val()
            };

            $.ajax({
                url: '/api/updateUser',  // 请根据后端API的实际路径修改
                type: 'POST',
                contentType: 'application/json',
                data: JSON.stringify(user),
                success: function () {
                    alert('用户信息修改成功');
                    $('#updateErrorInfo').html('');
                    // 如果需要，可以在此处刷新用户信息列表或进行其他操作
                },
                error: function (xhr, status, error) {
                    const errorMessage = xhr.responseJSON ? xhr.responseJSON.message : '用户信息修改失败';
                    $('#updateErrorInfo').html(errorMessage).css('color', 'red');
                }
            });
        });
    });

}

function getUserByUserName() {
    const username = $('#search').val();
    $.ajax({
        url: '/api/web10',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({ username: username }),
        success: function (data) {
            const table = $('#results');
            table.empty();
            table.append('<tr><th>ID</th><th>StudID</th><th>Username</th><th>Sex</th><th>Email</th></tr>');
            data.forEach(function (user) {
                table.append('<tr>' +
                    '<td>' + user.ID + '</td>' +
                    '<td>' + user.StudID + '</td>' +
                    '<td>' + user.Username + '</td>' +
                    '<td>' + user.Sex + '</td>' +
                    '<td>' + user.Email + '</td>' +
                    '</tr>');
            });
        },
        error: function (xhr, status, error) {
            alert('查询失败: ' + error);
        }
    });
}

function addUser() {
    const user = {
        StudID: $('#studID').val(),
        Username: $('#username').val(),
        Sex: $('#sex').val(),
        Email: $('#email').val(),
        Password: $('#password').val(),
    };
    $.ajax({
        url: '/api/register',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(user),
        success: function () {
            alert('用户添加成功');
            getUserByUserName()
        },
        error: function (xhr, status, error) {
            alert('用户添加失败: ' + error);
        }
    });
}

function delUser() {
    const id = $('#deleteUser').val();
    $.ajax({
        url: '/api/deleteUser',
        type: 'POST',
        contentType: 'text/plain', // 设置 contentType 为 text/plain
        data: id, // 直接发送 id 文本
        success: function () {
            console.log("删除成功");
            getUserByUserName()
        },
        error: function (xhr, status, error) {
            alert('删除失败: ' + error);
        }
    });
}

