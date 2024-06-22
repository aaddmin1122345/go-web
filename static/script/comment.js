$(document).ready(function() {
    // 提交评论表单
    $("#commentForm").on("submit", createComment);

    // 回复评论按钮点击事件
    $("#comments-section").on("click", "button", function() {
        const commentID = $(this).data("comment-id");
        replyToComment(commentID);
    });
});

function createComment(event) {
    event.preventDefault();

    const username = $("input[name='username']").val();
    const content = $(".content1").text().trim(); // 使用 .text() 获取文本内容并去除空格
    const currentUrl = window.location.href;
    const id = extractIdFromUrl(currentUrl);
    const parentCommentID = $("input[name='parentCommentID']").val() || null;

    const commentData = {
        username: username,
        content: content,
        articleID: id,
        parentCommentID: parseInt(parentCommentID)
    };

    console.log("Comment Data:", JSON.stringify(commentData));

    $.ajax({
        type: "POST",
        url: "/api/createComment",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(commentData),
        success: function () {
            alert("评论提交成功");
            location.reload(); // 成功提交后刷新页面
        },
        error: function (err) {
            let errorMessage = "评论提交失败";
            if (err.responseJSON && err.responseJSON.message) {
                errorMessage = err.responseJSON.message;
            }
            $("#errorInfo").html(errorMessage).css("color", "red"); // 显示错误消息给用户
        }
    });
}

function replyToComment(commentID) {
    $("input[name='parentCommentID']").val(commentID);
    $(".content1").focus(); // 设置焦点在评论内容输入框上
}

function extractIdFromUrl(url) {
    const regex = /[?&]id=(\d+)/;
    const match = regex.exec(url);
    return match ? parseInt(match[1], 10) : null;
}
