// pages/user_info/user_info.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        uname: '',
        sex: '',
        birthday: '',
        info: '',
        address: ''
    },

    getData:function(){
        let _this = this
        wx.cloud.callContainer({
            "config": {
              "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/userAuthenticationSee",
            "header": {
              "X-WX-SERVICE": "test-allsmile",
              "content-type": "application/json",
              "uid": app.globalData.uid
            },
            "method": "POST",
            "data": "",
            success: function (res) {
                let data = res.data.data
                _this.setData({
                    uname: data.uname,
                    sex: data.gender,
                    birthday: data.birthday,
                    info: data.resume,
                    address: data.block_address
                })
            }
          })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.getData()
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
        this.getData()
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