// pages/gainer_warehouse3/user_warehouse3.js

let app = getApp()

Page({

    /**
     * 页面的初始数据
     */
    data: {
        user: '',
        medical: '',
        m1: '',
        m2: ''
    },

    getData:function(){
        let _this = this
        wx.cloud.callContainer({
            "config": {
              "env": "prod-9gy59jvo10e0946b"
            },
            "path": "/gainer/DisplayWarehouse",
            "header": {
              "X-WX-SERVICE": "test-allsmile",
              "content-type": "application/json",
              "gid": app.globalData.gid
            },
            "method": "POST",
            "data": {
              "user": _this.data.user,
              "medical": _this.data.medical
            },
            success:function(res){
                _this.setData({
                    m1: res.data.data.M1,
                    m2: res.data.data.M2
                })
            }
          })
    },

    preview:function(e){
        let pic = []
        pic[0] = e.currentTarget.dataset.medical
        wx.previewImage({
          urls: pic,
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
            user: options.user
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