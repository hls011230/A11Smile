// pages/user_blockmess/user_blockmess.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        block_mess: []
    },

    getData: function () {
        let _this = this
        wx.showToast({
            title: '加载区块信息',
            icon: 'loading',
            duration: 5000
        })
        wx.request({
            url: 'https://test-allsmile-1687181-1310014865.ap-shanghai.run.tcloudbase.com/user/queryBlockInformation',
            timeout: 60000,
            method: 'POST',
            success: function (res) {
                let data = res.data.data
                _this.setData({
                    block_mess: data
                })
                wx.showToast({
                    title: '加载成功',
                    icon: 'success'
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