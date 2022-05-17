// pages/user_topup/user_topup.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        eth: '',
    },

    getEth: function (e) {
        this.setData({
            eth: e.detail.value
        })
    },

    ethForAs:function(){
        let _this = this
        wx.cloud.callContainer({
            "config": {
              "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/UETHforAs",
            "header": {
              "X-WX-SERVICE": "test-allsmile",
              "content-type": "application/json",
              "uid": app.globalData.uid
            },
            "method": "POST",
            "data": {
              "quantity": parseInt(_this.data.eth) 
            },
            success:function(res){
                wx.showToast({
                  title: '兑换成功',
                  icon: 'success',
                  duration: 2000
                })
            }
          })
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