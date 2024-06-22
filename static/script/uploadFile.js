    document.querySelector('.content').addEventListener('paste', function (event) {
    const items = (event.clipboardData || window.clipboardData).items;
    for (let item of items) {
    if (item.type.indexOf('image') !== -1) {
    const file = item.getAsFile();
    uploadImage(file);
}
}
});

    function uploadImage(file) {
    const formData = new FormData();
    formData.append('file', file);

    $.ajax({
    url: '/api/upload',
    type: 'POST',
    data: formData,
    processData: false,
    contentType: false,
    success: function (response) {
    const imageUrl = response.path;
    insertImageToEditor(imageUrl);
},
    error: function (jqXHR, textStatus, errorThrown) {
    console.error('Error uploading image:', textStatus, errorThrown);
}
});
}

    function insertImageToEditor(url) {
    const img = document.createElement('img');
    img.src = "/" + url; // 设置图片的 src 属性
    img.style.maxWidth = '80%'; // 设置图片样式以适应编辑器
    document.querySelector('.content').appendChild(img); // 将图片插入到编辑器中
}

    $(document).on("click", "#publishButton", function () {
    createArticle();
});