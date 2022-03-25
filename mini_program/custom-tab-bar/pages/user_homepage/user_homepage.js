Component({
    /**
     * 组件的属性列表
     */
    properties: {
  
    },
  
    /**
     * 组件的初始数据
     */
    data: {
      home:[{
        pagePath: "/pages/user_homepage/user_homepage",
        text: "首页",
        iconPath: "/pages/images/202203141622511.png",
        selectedIconPath: "/pages/images/20220314162251.png"
      },
      {
        pagePath: "/pages/user_data/user_data",
        text: "数据管理",
        iconPath: "/pages/images/916920790148990582.png",
        selectedIconPath: "/pages/images/169545831883691004.png"
      },
      {
        pagePath: "/pages/user_my/user_my",
        text: "我的",
        iconPath: "/pages/images/202203141622512.png",
        selectedIconPath: "/pages/images/202203141622513.png"
      }]
    },
  
    /**
     * 组件的方法列表
     */
    methods: {
  
    },
    pageLifetimes: {
      show() {
        if (typeof this.getTabBar === 'function' &&
          this.getTabBar()) {
          this.getTabBar().setData({
            list: this.data.home,
            selected: 0
          })
        }
      }
    }
  })
  