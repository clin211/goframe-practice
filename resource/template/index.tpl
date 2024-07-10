<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <style>
        .description {
            display: -webkit-box;
            -webkit-line-clamp: 3;
            -webkit-box-orient: vertical;
            overflow: hidden;
            text-overflow: ellipsis;
            height: 3rem;
            /* 控制描述的高度 */
            transition: height 0.3s ease;
        }
    </style>
    <script src="https://cdn.tailwindcss.com"></script>
</head>

<body>
    <div class="container mx-auto p-4">
        <h1
            class="relative pl-4 text-4xl font-bold mb-8 before:absolute before:inset-0 before:content before:block before:w-2 before:h-10 before:bg-red-500">
            Products</h1>
        <div class=" grid grid-cols-1 md:grid-cols-2 lg:grid-cols-5 gap-8">
            {{range .data}}
            <div class="bg-white rounded-lg shadow-lg">
                <img class="w-full h-64 p-2 object-cover mb-4 rounded" src="{{.Image}}" alt="{{.Title}}">
                <div class="p-6 pt-0">
                    <h2 class="text-xl font-bold mb-2 h-14 overflow-hidden">{{.Title}}</h2>
                    <p class="description text-gray-600 mb-2">{{.Description}}</p>
                    <div class="flex justify-between items-center">
                        <p class="text-gray-600 mb-2 text-red-500">${{.Price}}</p>
                        <p class="text-gray-600 mb-2">{{.Category}}</p>
                    </div>
                </div>
            </div>
            {{end}}
        </div>
    </div>
</body>

</html>