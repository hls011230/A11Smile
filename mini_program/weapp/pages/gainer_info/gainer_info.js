// pages/gainer_info/gainer_info.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        ename: '',
        resume: '',
        protrait: ''
    },

    getMess: function () {
        let _this = this
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/gainer/gainerAuthenticationSee",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "gid": app.globalData.gid
            },
            "method": "POST",
            "data": "",
            success: function (res) {
                let data = res.data.data
                _this.setData({
                    ename: data.enterprise_name,
                    info: data.resume
                })
            }
        })
    },

    getPortrait:function(){
        let _this = this
        wx.cloud.callContainer({
            "config": {
              "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/gainer/showGainerIcon",
            "header": {
              "X-WX-SERVICE": "test-allsmile",
              "content-type": "application/json",
              "gid": app.globalData.gid
            },
            "method": "POST",
            "data": "",
            success:function(res){
                if (res.data.code == 0) {
                    _this.setData({
                        protrait: res.data.data
                    })
                }else{
                    wx.showToast({
                      title: '服务器错误',
                      icon: 'error',
                      duration: '2000'
                    })
                }
            }
          })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.getMess()
        this.getPortrait()
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
        this.getMess()
        this.getPortrait()
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