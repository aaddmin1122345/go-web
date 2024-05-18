const user = {
    ID: 0,
    StudID: 0,
    Username: "",
    Sex: "",
    Email: "",
    Password: "",
    CreateTime: ""
};

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


function updateUser(userId) {
    // 获取更新后的用户信息
    const updatedUser = {
        StudID: $('#updateStudID_' + userId).val(),
        Username: $('#updateUsername_' + userId).val(),
        Sex: $('#updateSex_' + userId).val(),
        Email: $('#updateEmail_' + userId).val(),
        Password: $('#updatePassword_' + userId).val(),
    };

    $.ajax({
        url: '/api/updateUser',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(updatedUser),
        success: function () {
            alert('用户信息修改成功');
            getUserByUserName(); // 重新加载用户列表
        },
        error: function (xhr, status, error) {
            alert('用户信息修改失败: ' + error);
        }
    });
}


function getUserByUserName() {
    const username = $('#search').val();
    $.ajax({
        url: '/api/getUserByUserName',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({ username: username }),
        success: function (data) {
            const table = $('#results');
            table.empty();
            table.append('<tr><th>ID</th><th>StudID</th><th>Username</th><th>Sex</th><th>Email</th><th>Password</th><th>CreateTime</th><th>操作</th></tr>');
            data.forEach(function (user) {
                const tr = $('<tr>');
                tr.append('<td>' + user.ID + '</td>');
                tr.append('<td>' + user.StudID + '</td>');
                tr.append('<td id="username_' + user.ID + '">' + user.Username + '</td>');
                tr.append('<td id="sex_' + user.ID + '">' + user.Sex + '</td>');
                tr.append('<td id="email_' + user.ID + '">' + user.Email + '</td>');
                tr.append('<td id="password_' + user.ID + '">' + user.Password + '</td>');
                tr.append('<td>' + user.CreateTime + '</td>');

                // 添加编辑按钮
                const editButton = $('<button>').text('编辑');
                editButton.click(function () {
                    const userId = user.ID; // 获取用户的 ID

                    // 将该行信息显示为输入框
                    $('#username_' + userId).html('<input type="text" id="updateUsername_' + userId + '" value="' + user.Username + '">');
                    $('#sex_' + userId).html('<input type="text" id="updateSex_' + userId + '" value="' + user.Sex + '">');
                    $('#email_' + userId).html('<input type="text" id="updateEmail_' + userId + '" value="' + user.Email + '">');
                    $('#password_' + userId).html('<input type="text" id="updatePassword_' + userId + '" value="' + user.Password + '">');

                    // 替换编辑按钮为更新按钮
                    const updateButton = $('<button>').text('更新');
                    updateButton.click(function () {
                        updateUser(userId); // 传递用户的 ID 给 updateUser 函数
                    });
                    $(this).replaceWith(updateButton);

                    // 取消删除按钮的点击事件
                    $(this).siblings('td').find('button:contains("删除")').off('click');
                });
                tr.append($('<td>').append(editButton));

                // 添加删除按钮
                const deleteButton = $('<button>').text('删除');
                deleteButton.click(function () {
                    delUser(user.ID);
                });
                tr.append($('<td>').append(deleteButton));

                table.append(tr);
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
            getUserByUserName(); // 重新加载用户列表
        },
        error: function (xhr, status, error) {
            alert('用户添加失败: ' + error);
        }
    });
}

function delUser(userId) {
    $.ajax({
        url: '/api/deleteUser',
        type: 'POST',
        contentType: 'text/plain',
        data: userId.toString(),
        success: function () {
            console.log("删除成功");
            getUserByUserName();
        },
        error: function (xhr, status, error) {
            alert('删除失败: ' + error);
        }
    });
    console.log(typeof(userId))
}

