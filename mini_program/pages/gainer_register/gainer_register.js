// pages/user_register/user_register.js
const date = new Date();
var year = date.getFullYear();
var month = date.getMonth() + 1;
var day = date.getDate();
if (month < 10) {undefined
month = "0" + month;
}
if (day < 10) {undefined
day = "0" + day;
}
var nowDate = year + "-" + month + "-" + day;
Page({

    /**
     * 页面的初始数据
     */
    data: {
      sexArr : ['男','女'],
      nowDate : nowDate
    },

  bindDateChange: function (e) {
    console.log(e.detail.value)
   this.setData({
     dates: e.detail.value
   })
  },

  bindSexChange: function (e) {
    this.setData({
      sex : e.detail.value
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