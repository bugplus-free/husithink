<!DOCTYPE html>
<html lang="zh">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>网站起始页</title>
    <style>
        /* 基础样式 */
        body {
            margin: 0;
            overflow-x: hidden;
        }

        .carousel-container {
            position: fixed; /* 修改为fixed定位 */
            top: 0; /* 添加top: 0使容器顶部紧贴屏幕顶部 */
            left: 0; /* 添加left: 0使容器左侧紧贴屏幕左侧 */
            width: 100%; /* 修改宽度为100% */
            height: 100vh; /* 修改高度为100vh，表示占据整个视口高度 */
            overflow: hidden;
            display: flex;
            justify-content: center;
            align-items: center;
        }

        .carousel-item {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            object-fit: cover;
            opacity: 0;
            transition: opacity 1s ease-in-out;
            display: flex;
            justify-content: center;
            align-items: center;
        }

        .carousel-item img {
            max-width: 100%;
            max-height: 100%;
            object-fit: contain;
            display: block;
        }

        .carousel-item.active {
            opacity: 1;
        }

        .login-register {
            position: fixed;
            top: 24px;
            right: 24px;
            z-index: 1000;
        }

        .login-register button {
            display: inline-block;
            padding: 12px 16px;
            margin-left: 8px;
            background-color: #007bff;
            color: #fff;
            font-size: 16px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            transition: background-color 0.2s ease-in-out;
        }

        .login-register button:hover {
            background-color: #0069d9;
        }
    </style>

    <!-- 走马灯轮播效果（仅提供基本逻辑，实际应用可能需要使用JavaScript库如Swiper实现更复杂的功能） -->
    <script>
    
        function carousel() {
            const container = document.querySelector('.carousel-container');
            const items = container.querySelectorAll('.carousel-item');
            let activeIndex = 0;

            function showNextItem() {
                items[activeIndex].classList.remove('active');
                activeIndex = (activeIndex + 1) % items.length;
                items[activeIndex].classList.add('active');
            }

            setInterval(showNextItem, 3000); // 每5秒切换一张图片

            showNextItem(); // 初始显示第一张图片
        }

        document.addEventListener('DOMContentLoaded', carousel);



        function generateToken(salt) {
            const hash = new Uint8Array(16); // 创建一个16字节的Uint8Array，用于存放MD5哈希结果
            const now = new Date().toISOString(); // 获取当前时间字符串

            // 在浏览器环境中，可以使用Web Cryptography API来计算MD5哈希
            window.crypto.subtle.digest('md5', new TextEncoder().encode(salt + now))
                .then(hash => {
                const token = Array.from(new Uint8Array(hash)).map(b => b.toString(16).padStart(2, '0')).join('');
                return token;
                })
                .catch(err => {
                console.error('Error generating token:', err);
                return null; // 或者抛出错误，根据您的需求处理
                });
            }

            // 使用示例
            generateToken("astaxie%^7&8888").then(token => {
            // 发送请求时附带生成的token
            fetch('/your-api-endpoint', {
                method: 'POST',
                headers: {
                'Content-Type': 'application/json',
                'X-Token': token, // 将token放在自定义请求头中，具体名称根据您的服务端要求调整
                },
                body: JSON.stringify({ /* 请求体数据 */ }),
            })
                .then(response => {
                // 处理响应
                })
                .catch(error => {
                // 处理请求错误
                });
            });
    </script>
</head>
<body>
    <div class="carousel-container">
        <img class="carousel-item active" src="/src/images/begin1.jpg" alt="Image 1">
        <img class="carousel-item" src="/src/images/begin2.jpg" alt="Image 2">
        <img class="carousel-item" src="/src/images/begin3.jpg" alt="Image 3">
        <img class="carousel-item" src="/src/images/begin0.jpg" alt="Image 4">
    </div>

    <div class="login-register">
        <button onclick="location.href='/login'">登录</button>
        <button onclick="location.href='/enroll'">注册</button>
    </div>
</body>
</html>