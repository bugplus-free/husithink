<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>用户登录与注册</title>
    <style>
        /* 基础样式 */
        body {
            font-family: Arial, sans-serif;
            background-color: #f5f5f5;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
            padding: 0;
        }

        .container {
            background-color: #fff;
            box-shadow: 0px 4px 16px rgba(0, 0, 0, 0.1);
            border-radius: 4px;
            width: 320px;
            padding: 32px;
        }

        h1 {
            text-align: center;
            margin-bottom: 24px;
        }

        form {
            display: flex;
            flex-direction: column;
        }

        input[type="text"],
        input[type="password"] {
            width: 100%;
            padding: 12px 16px;
            margin-bottom: 16px;
            border: 1px solid #ccc;
            border-radius: 4px;
            font-size: 16px;
            box-sizing: border-box;
        }

        input[type="submit"] {
            width: 100%;
            padding: 12px 16px;
            background-color: #007bff;
            color: #fff;
            font-size: 16px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            transition: background-color 0.2s ease-in-out;
        }

        input[type="submit"]:hover {
            background-color: #0069d9;
        }

        .link-button {
            display: inline-block;
            padding: 12px 16px;
            background-color: transparent;
            color: #007bff;
            font-size: 16px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            text-decoration: none;
            transition: color 0.2s ease-in-out;
        }

        .link-button:hover {
            color: #0069d9;
        }

        .separator {
            margin-top: 24px;
            margin-bottom: 24px;
            text-align: center;
            font-size: 14px;
            color: #6c757d;
        }

        .icon {
            width: 24px;
            height: 24px;
            margin-right: 8px;
            vertical-align: middle;
        }
    </style>
    <script>
        function validateUsername(input) {
            const regex = /^[A-Za-z]+$/;
            return regex.test(input.value);
        }

        function validatePassword(input) {
            const regex = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,16}$/;
            return regex.test(input.value);
        }

        function validateInput(input, validator, errorMessage) {
            const isValid = validator(input);
            if (!isValid) {
                input.setCustomValidity(errorMessage);
            } else {
                input.setCustomValidity('');
            }
        }
        document.addEventListener('DOMContentLoaded', () => {
            const usernameInput = document.querySelector('input[name="username"]');
            const passwordInput = document.querySelector('input[name="password"]');

            usernameInput.addEventListener('input', () => {
                validateInput(usernameInput, validateUsername, '用户名只能包含英文大小写字母');
            });

            passwordInput.addEventListener('input', () => {
                validateInput(passwordInput, validatePassword, '密码必须是8到16位的英文大小写字母和数字组合');
            });
        });

        

        // 使用示例
        
    </script>
</head>
<body>
    <div class="container">
        <h1>欢迎登录</h1>
        <form action="/login" method="post" id="login-form">
            <div class="input-group">
                <input type="text" name="username" placeholder="用户名" required>
            </div>
            <div class="input-group">
                <input type="password" name="password" placeholder="密码" required>
            </div>
            <input type="hidden" name="token" value="{{.}}">
            <input type="submit" value="登录">
        </form>
        <div class="separator">或</div>
        <a href="/enroll" class="link-button">注册新账户</a>
    </div>
</body>
</html>