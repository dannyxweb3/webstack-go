      // 示例数据（实际可从后端API获取） // 只推荐3，4个
      let featuredRecommendations = {}; // #featuredRecommendations

      // 随机选择N个推荐项
      function getFeaturedRecommendations(count = 3) {
          const shuffled = [...featuredRecommendations].sort(() => 0.5 - Math.random());
          return shuffled.slice(0, count);
      }

      // 渲染推荐列表
      function renderFeaturedRecommendations() {
          const recommendations = getFeaturedRecommendations();
          const container = document.getElementById('featured-recommendList');
          container.innerHTML = '';
          
          recommendations.forEach(item => {
              container.innerHTML += `
                  <a href="${item.url}" target="_blank" class="recommend-item">
                      <div class="recommend-header">
                          <img src="${item.logo}" class="recommend-logo" alt="${item.name}" 
                               __onerror="this.src='https://via.placeholder.com/32'">
                          <div class="recommend-name">${item.name}</div>
                      </div>
                      <div class="recommend-desc">${item.desc}</div>
                  </a>
              `;
          });
      }

      // 初始化加载
      $(document).ready(function() {
          featuredRecommendations = JSON.parse($('#featuredRecommendations').val());
          renderFeaturedRecommendations();
          
          // 点击刷新按钮换一批
          $('#featured-refreshBtn').off('click').on('click', function() {
              refreshFeaturedRecommendations();
          });
      });

      // 暴露刷新方法供其他部分调用
      function refreshFeaturedRecommendations() {
        renderFeaturedRecommendations();
      }
