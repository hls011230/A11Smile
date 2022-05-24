// pages/gainer_warehouse2/user_warehouse2.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        address: '',
        medical: '',
        users: []
    },

    getData:function(){
        let _this = this
        wx.cloud.callContainer({
            "config": {
              "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/gainer/DisplayWarehouseUser",
            "header": {
              "X-WX-SERVICE": "test-allsmile",
              "content-type": "application/json",
              "gid": app.globalData.gid
            },
            "method": "POST",
            "data": {
              "soliciter": _this.data.address,
              "medical": _this.data.medical
            },
            success:function(res){
                _this.setData({
                    users: res.data.data
                })
            }
          })
    },

    getCertificate:function(e){
        let user = e.currentTarget.dataset.user
        let medical = e.currentTarget.dataset.medical
        wx.navigateTo({
          url: '/pages/gainer_warehouse3/user_warehouse3?user='+user+"&medical="+medical,
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        this.setData({
            medical: options.medical
        })

        this.setData({
            address: options.address
        })

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