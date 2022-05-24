// pages/user_classify/user_classify.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        sid: '',
        arr: []
    },

    goDetailsPage: function (e) {
        let medical_name = e.currentTarget.dataset.medical_name
        let address = e.currentTarget.dataset.address
        let hospital_name = e.currentTarget.dataset.hospital_name
        wx.navigateTo({
            url: "/pages/user_uploading/user_uploading?medical_name="+medical_name+"&address="+address+"&hospital_name="+hospital_name
        })
    },

    classify: function (e) {
        let _this = this
        wx.request({
            url: 'https://test-allsmile-1687181-1310014865.ap-shanghai.run.tcloudbase.com/user/showSortPage?sid=' + _this.data.sid,
            method: "POST",
            success: function (res) {
                _this.setData({
                    arr: res.data.data
                })
            }
        })
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        this.setData({
            sid: options.classify
        })
        this.classify()
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