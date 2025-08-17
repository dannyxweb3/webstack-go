      // 侧边栏热门分类 从隐藏元素加载分类数据
      function loadCategoriesData() {
        const categories = [
            {
                id: 1,
                name: "Free AI Tools",
                url: "/free-ai-tools"
            },
            {
                id: 2,
                name: "AI Image Generation",
                url: "/category/image-generation-editing"
            },
            {
                id: 3,
                name: "AI Text Generator",
                url: "/category/ai-text-generator"
            },
            {
                id: 4,
                name: "AI Scipt To Video",
                url: "/category/ai-video-editor"
            },
            {
                id: 5,
                name: "AI Money Making",
                url: "/category/business-research"
            }
        ];
        $('#categoriesData > div').each(function() {
            categories.push({
                id: $(this).data('id'),
                name: $(this).data('name'),
                url: $(this).data('url')
            });
        });
        return categories;
    }

    // 渲染分类标签
    function renderCategories() {
        const categories = loadCategoriesData();
        const container = document.getElementById('categoriesList');
        container.innerHTML = '';
        
        categories.forEach(category => {
            container.innerHTML += `
                <a href="${category.url}" class="category-tag" 
                  data-id="${category.id}">
                    ${category.name}
                </a>
            `;
        });
    }

    // 页面加载时初始化
    $(document).ready(function() {
        renderCategories();
    });
