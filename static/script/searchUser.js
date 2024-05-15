"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const axios_1 = __importDefault(require("axios"));
function searchUser() {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const searchUser = document.getElementById("searchUser").value;
            console.log(searchUser);
            const response = yield axios_1.default.post("/api/searchUsers", { searchUser });
            if (response.status !== 200) {
                throw new Error("Failed to fetch data. Status: " + response.status);
            }
            const data = response.data;
            console.log("success", data);
            document.getElementById("aaa").innerHTML = data;
            document.getElementById("aaa").style.color = "red";
            // 处理成功响应的逻辑
        }
        catch (error) {
            console.error("Error:", error.message); // 添加类型断言
            // 处理失败响应的逻辑
        }
    });
}
