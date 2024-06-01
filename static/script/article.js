function createArticle() {
    const articleData = {
        title: $("#title").val(),
        content: $("#content").val(),
        imageURL: "",
        category: $("#category").val(),
    };

    $.ajax({
        type: "POST",
        url: "/api/createArticle",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(articleData),
        success: function (data) {
            alert("文章发布成功");
            window.location.href = "/";
        },
        error: function (err) {
            let errorMessage = "文章发布失败";
            if (err.responseJSON && err.responseJSON.message) {
                errorMessage = err.responseJSON.message;
            }
            $("#errorInfo").html(errorMessage).css("color", "red");
        }
    });
}

$(document).ready(function () {
    $("#articleForm").submit(function (event) {
        event.preventDefault(); // 阻止表单默认提交行为
        createArticle();
    });
});
