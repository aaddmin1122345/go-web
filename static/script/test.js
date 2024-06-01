function getArticleByKeyword() {
    const keyword = $('#search').val();
    $.ajax({
        url: '/api/getArticleByKeyword',
        type: 'POST',
        contentType: 'application/json',
        data: JSON.stringify({keyword: keyword}),
        success: function (data) {
            const table = $('#results');
            table.find("tr:gt(0)").remove(); // Clear all rows except the first (header row)

            data.forEach(function (article) {
                const tr = $('<tr>');
                tr.append('<td>' + article.ID + '</td>');
                tr.append('<td>' + article.Title + '</td>');
                tr.append('<td>' + article.Content + '</td>');
                tr.append('<td>' + article.CreateDate + '</td>');
                tr.append('<td>' + article.ImageURL + '</td>');
                tr.append('<td>' + article.Category + '</td>');

                // 编辑按钮
                const editButton = $('<button>').text('编辑').click(function () {
                    editArticle(article);
                });
                tr.append($('<td>').append(editButton));

                // 删除按钮
                const deleteButton = $('<button>').text('删除').click(function () {
                    delArticle(article.ID);
                });
                tr.append($('<td>').append(deleteButton));

                table.append(tr);
            });
        },
        error: function (xhr, status, error) {
            console.log('查询失败: ' + error);
        }
    });
}

function editArticle(article) {
    // Define the logic for editing an article
    console.log('Edit article:', article);
    // Implement the edit functionality here
}

function delArticle(articleId) {
    // Define the logic for deleting an article
    console.log('Delete article ID:', articleId);
    // Implement the delete functionality here
}
