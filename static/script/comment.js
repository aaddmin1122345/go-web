function createComment(event) {
    event.preventDefault(); // Prevent default form submission

    // Get username and content values
    const username = $("input[name='username']").val();
    const content = $("textarea[name='content']").val();

    // Get the current URL and extract 'id' parameter
    const currentUrl = window.location.href;
    const id = parseInt(extractIdFromUrl(currentUrl), 10); // Parse 'id' as integer

    // Construct comment data
    const commentData = {
        username: username,
        content: content,
        articleID: id
    };

    // Output JSON data to console (for debugging)
    console.log("Comment Data:", JSON.stringify(commentData));

    // Send comment data to the backend using AJAX
    $.ajax({
        type: "POST",
        url: "/api/createComment",
        contentType: "application/json; charset=utf-8",
        data: JSON.stringify(commentData),
        success: function () {
            alert("评论成功");
            location.reload(); // Reload the page after successful comment submission
        },
        error: function (err) {
            let errorMessage = "发送失败";
            if (err.responseJSON && err.responseJSON.message) {
                errorMessage = err.responseJSON.message;
            }
            $("#errorInfo").html(errorMessage).css("color", "red"); // Display error message to user
        }
    });
}

// Function to extract 'id' parameter from URL using regex
function extractIdFromUrl(url) {
    // Regular expression to match 'id' parameter in URL
    const regex = /[?&]id=(\d+)/;
    const match = regex.exec(url);
    return match && match[1]; // Return the captured group (id value)
}

// Attach the createComment function to the form's submit event
$(document).ready(function() {
    $("#commentForm").on("submit", createComment);
});
