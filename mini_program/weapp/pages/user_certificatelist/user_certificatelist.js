// pages/user_certificatelist/user_certificatelist.js\

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        certificatelist: [],
        serial: ''
    },

    getData: function () {
        let _this = this
        wx.cloud.callContainer({
            "config": {
                "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/user/showAllCertificate",
            "header": {
                "X-WX-SERVICE": "test-allsmile",
                "content-type": "application/json",
                "uid": app.globalData.uid
            },
            "method": "POST",
            "data": {},
            success:function(res){
                _this.setData({
                    certificatelist: res.data.data
                })
            }
        })
    },

    showDetailsCertificate:function(e){
        let _this = this
        let serial = e.currentTarget.dataset.serial
        _this.setData({
           serial: serial 
        })
        wx.navigateTo({
            url: '/pages/user_certificate/user_certificate',
            success: function (res) {
                res.eventChannel.emit('serial', { data: _this.data.serial})
            }
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        this.getData()
    },

    /**
     * 生命周期函数--监听页面初次渲染完成
     */
    onReady() {

    },

    /**
     * 生命周期函数--监听页面显示
     */
    onShow() {
        this.getData()
    },

    /**
     * 生命周期函数--监听页面隐藏
     */
    onHide() {

    },

    /**
     * 生命周期函数--监听页面卸载
     */
    onUnload() {

    },

    /**
     * 页面相关事件处理函数--监听用户下拉动作
     */
    onPullDownRefresh() {

    },

    /**
     * 页面上拉触底事件的处理函数
     */
    onReachBottom() {

    },

    /**
     * 用户点击右上角分享
     */
    onShareAppMessage() {

    }
})