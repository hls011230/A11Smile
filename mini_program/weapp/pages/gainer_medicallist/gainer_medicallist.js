// pages/gainer_medicallist/gainer_medicallist.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        medicals: [],
        address: ''
    },

    getData:function(){
        let _this = this
        wx.cloud.callContainer({
            "config": {
              "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/gainer/viewMedicalName",
            "header": {
              "X-WX-SERVICE": "test-allsmile",
              "content-type": "application/json",
              "gid": app.globalData.gid
            },
            "method": "POST",
            "data": "",
            success: function(res){
                _this.setData({
                    medicals: res.data.data
                })
            }
          })
    },

    getAdd:function(){
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
            success: function(res){
                _this.setData({
                    address: res.data.data.block_address
                })
            }
          })
    },

    getUser:function(e){
        let address = e.currentTarget.dataset.address
        let medical = e.currentTarget.dataset.medical
        wx.navigateTo({
          url: '/pages/gainer_warehouse2/user_warehouse2?address='+address+"&medical="+medical,
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad: function (options) {
        this.getAdd()
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
        this.getAdd()
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