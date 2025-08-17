      // 从隐藏元素加载新闻数据
      function loadNewsData() {
        const newsItems = [
            {
                time: "2025-01-01 10:00",
                title: "The OpenAI News page provides updates on the company’s latest research",
                logo: "https://openai.com/favicon.ico",
                source: "OpenAI News",
                src: "https://openai.com/news",
            },
            {
                time: "2025-01-02 12:30",
                title: "Google Research Blog provides in-depth articles, insights, and updates on the latest innovations, projects",
                logo: "https://research.google/favicon.ico",
                source: "AI Daily",
                src: "https://research.google/blog"
            },
            {
                time: "2025-01-03 15:45",
                title: "The AI blog by Meta shares updates, insights, and research breakthroughs from Meta’s AI teams, focusing on advancements in artificial intelligence",
                logo: "https://ai.meta.com/favicon.ico",
                source: "Meta AI Blog",
                src: "https://ai.meta.com/blog"
            }
        ];
        $('#newsData > div').each(function() {
            newsItems.push({
                time: $(this).data('time'),
                title: $(this).data('title'),
                logo: $(this).data('logo'),
                source: $(this).data('source')
            });
        });
        return newsItems;
    }

    // 渲染新闻列表
    function renderNewsList() {
        const newsData = loadNewsData();
        const container = document.getElementById('newsList');
        container.innerHTML = '';
        
        newsData.forEach(item => {
            container.innerHTML += `
                <a href="${item.src}" class="news-item" target="_blank" alt="${item.title}">
                    <div class="news-header">
                        <img src="${item.logo}" class="news-logo" alt="${item.source}"
                            __onerror="this.src='https://via.placeholder.com/24'" style="display:none;">
                        <div class="news-title">${item.title}</div>
                    </div>
                    <div class="news-meta">
                        <span>${item.time}</span> · 
                        <span>${item.source}</span>
                    </div>
                </a>
            `;
        });
    }

    // 页面加载时初始化
    $(document).ready(function() {
        renderNewsList();
        
        // 点击新闻项跳转（示例，实际需绑定正确URL）
        $(document).on('click', '.news-item', function(e) {
            e.preventDefault();
            alert('实际项目中这里会跳转到新闻详情页');
            // window.open(newsUrl, '_blank');
        });
    });
