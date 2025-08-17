      // 示例数据（实际可从后端API获取）
      let randomlyRecommendations = []; // #randomlyRecommendations

      // 随机选择N个推荐项
      function getRandomRecommendations(count = 10) {
          const shuffled = [...randomlyRecommendations].sort(() => 0.5 - Math.random());
          return shuffled.slice(0, count);
      }

      // 渲染推荐列表
      function renderRecommendations() {
          const recommendations = getRandomRecommendations();
          const container = document.getElementById('randomly-recommendList');
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
          randomlyRecommendations = JSON.parse($('#randomlyRecommendations').val());
          renderRecommendations();
          
          // 点击刷新按钮换一批
          $('#randomly-refreshBtn').off('click').on('click', function() {
              refreshRandomlyRecommendations();
          });
      });

      // 暴露刷新方法供其他部分调用
      function refreshRandomlyRecommendations() {
          renderRecommendations();
      }
