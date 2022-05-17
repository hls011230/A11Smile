// pages/user_order/user_order.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        isTabs: 1,
        undoneOrderData: [],
        doneOrderData: []
    },

    checkTap(e) {
        this.setData({
            isTabs:e.currentTarget.dataset.flag
        })
    },

    getUndoneData:function(){
        let _this = this
        wx.cloud.callContainer({
            "config": {
              "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/NoTransactions",
            "header": {
              "X-WX-SERVICE": "test-allsmile",
              "content-type": "application/json",
              "uid": app.globalData.uid
            },
            "method": "POST",
            "data": "",
            success:function(res){
                _this.setData({
                    undoneOrderData: res.data.data
                })
            }
          })
    },

    getDoneData: function(){
        let _this = this
        wx.cloud.callContainer({
            "config": {
              "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/AllTransactions",
            "header": {
              "X-WX-SERVICE": "test-allsmile",
              "content-type": "application/json",
              "uid": app.globalData.uid
            },
            "method": "POST",
            "data": "",
            success:function(res){
                _this.setData({
                    doneOrderData: res.data.data
                })
            }
          })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.getUndoneData()
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
        this.getUndoneData()
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