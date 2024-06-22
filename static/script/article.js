function createArticle() {
    // 获取文章标题、内容和分类
    const title = $(".title").val();
    const content = $(".content").html();
    const category = $("input[name='category']:checked").val();
    const headURL = $(".content img").attr("src");

    // 构建文章数据
    const articleData = {
        title: title,
        content: content,
        imageURL: headURL,
        category: category,
        // userID : userID,
    };

    // 添加文件路径到文章数据（如果有上传的文件路径）
    if (uploadedFilePath) {
        articleData.file = uploadedFilePath;
    }

    // 输出发送给后端的 JSON 数据
    console.log("json", JSON.stringify(articleData));

    // 发送文章数据到后端
    $.ajax({
        type: "POST",
        url: "/api/createArticle",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(articleData),
        success: function () {
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
