// pages/gainer_my/gainer_my.js
Page({

    /**
     * 页面的初始数据
     */
    data: {

    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {

    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady: function () {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow: function () {

    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide: function () {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload: function () {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh: function () {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom: function () {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage: function () {

    }
})

Component({
    data: {
      index: [{
        pagePath: "/pages/gainer_homepage/gainer_homepage",
        iconPath: "/pages/images/202203141622511.png",
        selectedIconPath: "/pages/images/20220314162251.png",
        text: "首页"
      }, {
        pagePath: "/pages/gainer_my/gainer_my",
        iconPath: "/pages/images/202203141622512.png",
        selectedIconPath: "/pages/images/202203141622513.png",
        text: "我的"
      }]
    },
    pageLifetimes: {
      show() {
        if (typeof this.getTabBar === 'function' &&
          this.getTabBar()) {
          this.getTabBar().setData({
            list: this.data.index,
            selected: 1
          })
        }
      }
    }
  })