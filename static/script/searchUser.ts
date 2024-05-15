import axios from 'axios';

async function searchUser(): Promise<void> {
    try {
        const searchUser: string = (document.getElementById("searchUser") as HTMLInputElement).value;
        console.log(searchUser);

        const response = await axios.post("/api/searchUsers", { searchUser });

        if (response.status !== 200) {
            throw new Error("Failed to fetch data. Status: " + response.status);
        }

        const data = response.data;
        console.log("success", data);
        document.getElementById("aaa")!.innerHTML = data;
        document.getElementById("aaa")!.style.color = "red";
        // 处理成功响应的逻辑
    } catch (error) {
        console.error("Error:", (error as Error).message); // 添加类型断言
        // 处理失败响应的逻辑
    }
}

