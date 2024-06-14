function createArticle() {
    // 获取文章标题、内容和分类
    const title = $(".title").val();
    const content = $(".content").html();
    const category = $("input[name='category']:checked").val();
    const headURL = $(".content img").attr("src");
    // 获取页面上的所有图片
    // const images = [];
    // $(".content img").each(function() {
    //     images.push($(this).attr("src"));
    // });

    // 构建文章数据
    const articleData = {
        title: title,
        content: content,
        imageURL: headURL, // 将图片地址数组添加到文章数据中
        category: category,
    };

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
