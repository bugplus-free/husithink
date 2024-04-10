<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>用户注册</title>
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

        .input-group {
            position: relative;
            margin-bottom: 16px;
        }

        input[type="text"],
        input[type="password"] {
            width: 100%;
            padding: 12px 16px;
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

        .error-message {
            position: absolute;
            top: ¼em; /* 调整为合适的距离 */
            left: 0;
            width: 100%;
            font-size: 14px;
            color: #dc3545;
            text-align: left;
            visibility: hidden;
            transition: visibility 0s, opacity 0.3s linear;
        }

        .input-group.has-error .error-message {
            visibility: visible;
            opacity: 1;
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

        function validateEmail(input) {
            const regex = /^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$/;
            return regex.test(input.value);
        }

        function validatePassword(input) {
            const regex = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,16}$/;
            return regex.test(input.value);
        }

        function validateInput(input, validator, errorMessage) {
            const isValid = validator(input);
            const errorElement = input.nextElementSibling;

            if (!isValid) {
                errorElement.textContent = errorMessage;
                input.parentElement.classList.add('has-error');
            } else {
                errorElement.textContent = '';
                input.parentElement.classList.remove('has-error');
            }
        }
        document.addEventListener('DOMContentLoaded', () => {
            const usernameInput = document.querySelector('input[name="username"]');
            const emailInput = document.querySelector('input[name="email"]');
            const passwordInput = document.querySelector('input[name="password"]');

            usernameInput.addEventListener('input', () => {
                validateInput(usernameInput, validateUsername, '用户名只能包含英文大小写字母');
            });

            emailInput.addEventListener('input', () => {
                validateInput(emailInput, validateEmail, '请输入有效的电子邮件地址');
            });

            passwordInput.addEventListener('input', () => {
                validateInput(passwordInput, validatePassword, '密码必须是8到16位的英文大小写字母和数字组合');
            });
        });
    </script>
</head>
<body>
    <div class="container">
        <h1>用户注册</h1>
        <form action="/enroll" method="post">
            <div class="input-group">
                <input type="text" name="username" placeholder="用户名" required>
                <span class="error-message"></span>
            </div>
            <div class="input-group">
                <input type="text" name="email" placeholder="邮箱" required>
                <span class="error-message"></span>
            </div>
            <div class="input-group">
                <input type="password" name="password" placeholder="密码" required>
                <span class="error-message"></span>
            </div>
            <input type="submit" value="注册">
        </form>
    </div>
</body>
</html>