function updateUser(userId) {
    // 获取更新后的用户信息
    const updatedUser = {
        ID: userId,
        PhoneNum: $('#updatePhoneNum_' + userId).val(),
        Username: $('#updateUsername_' + userId).val(),
        Sex: $('#updateSex_' + userId).val(),
        Email: $('#updateEmail_' + userId).val(),
        Password: $('#updatePassword_' + userId).val(),
        UserType: $('#updateUserType_' + userId).val(),
    };

    $.ajax({
        url: '/api/updateUser',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify(updatedUser),
        success: function () {
            alert('用户信息修改成功');
            getUserByKeyword() // 重新加载用户列表
        },
    });
}


// 重新写的模糊查询
function getUserByKeyword() {
    const username = $('#search').val();
    $.ajax({
        url: '/api/getUserByKeyword',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({ username: username }),
        success: function (data) {
            const table = $('#results');
            table.empty();
            table.append('<tr><th>ID</th><th>手机号</th><th>用户名</th><th>性别</th><th>邮箱</th><th>密码</th><th>注册类型</th><th>创建时间</th><th>操作</th></tr>');
            data.forEach(function (user) {
                const tr = $('<tr>');
                tr.append('<td>' + user.ID + '</td>');
                tr.append('<td>' + user.PhoneNum + '</td>');
                tr.append('<td>' + user.Username + '</td>');
                tr.append('<td>' + user.Sex + '</td>');
                tr.append('<td>' + user.Email + '</td>');
                tr.append('<td>' + user.Password + '</td>');
                tr.append('<td>' + user.UserType + '</td>');
                tr.append('<td>' + user.CreateTime + '</td>');

                // 编辑按钮
                const editButton = $('<button>').text('编辑').click(function () {
                    editUser(user);
                });
                tr.append($('<td>').append(editButton));

                // 删除按钮
                const deleteButton = $('<button>').text('删除').click(function () {
                    delUser(user.ID);
                });
                tr.append($('<td>').append(deleteButton));

                table.append(tr);
            });
        },
        error: function (xhr, status, error) {
            // alert('查询失败: ' + error);
            console.log(error)
        }
    });
}


// 不加这个提交了会自动刷新
$(document).ready(function() {
    $('#registerForm').submit(function(event) {
        event.preventDefault(); // 防止默认表单提交行为
        addUser();
    });
});




function addUser() {
    const user = {
        PhoneNum: $('#phoneNum').val(),
        Username: $('#username').val(),
        Sex: $('#sex').val(),
        Email: $('#email').val(),
        Password: $('#password').val(),
        UserType: $('#userType').val(),
    };

    if (!user.PhoneNum || !user.Username || !user.Sex || !user.Email || !user.Password|| !user.UserType) {
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
            getUserByKeyword() // 重新加载用户列表
        },
        error: function (xhr, status, error) {
            console.log(error);
        }
    });
}


function delUser(userId) {
    $.ajax({
        url: '/api/delUser',
        type: 'POST',
        contentType: 'text/plain',
        data: userId.toString(),
        success: function () {
            console.log("删除成功");
            getUserByKeyword()
        },
        error: function (xhr, status, error) {
            alert('删除失败: ' + error);
        }
    });
    console.log(typeof (userId))
}


function editUser(user) {
    const userId = user.ID;
    // 加上 userId 的原因是确保在 HTML 文档中对特定用户的信息进行操作，而不会误操作其他用户的信息。
    $('#phoneNum_' + userId).html('<input type="text" id="updatePhoneNum_' + userId + '" value="' + user.PhoneNum + '">');
    $('#username_' + userId).html('<input type="text" id="updateUsername_' + userId + '" value="' + user.Username + '">');
    $('#sex_' + userId).html('<input type="text" id="updateSex_' + userId + '" value="' + user.Sex + '">');
    $('#email_' + userId).html('<input type="text" id="updateEmail_' + userId + '" value="' + user.Email + '">');
    $('#password_' + userId).html('<input type="text" id="updatePassword_' + userId + '" value="' + user.Password + '">');
    $('#userType_' + userId).html('<input type="text" id="updateUserType_' + userId + '" value="' + user.UserType + '">');


    const updateButton = $('<button>').text('更新').click(function () {
        updateUser(userId);
    });
    // $('#phoneNum_' + userId).closest('tr').find('button:contains("编辑")').replaceWith(updateButton);

    $('#results').find('button:contains("编辑")').replaceWith(updateButton);
}


