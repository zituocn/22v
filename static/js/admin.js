
function autotitle(){


}

function checkdown(){
	var val = getValue("hdtv");
	//alert(val);

	return true;
}
function getValue(name){
	var array = new Array();
	$("input[name='"+name+"']").each(function(){
		var val = $(this).val();
		if(val.length>0)
			array.push($(this).val());
	});
	return array.join(",");
}

function checkpost(){
	var cid = $("#cid");
	var name  = $("#name");
	var ename = $("#ename");
	var episode = $("#episode");
	var actor = $("#actor");
	var director = $("#director");
	var updateweek = $("#updateweek");
	var playdate = $("#playdate");
	var content = $("#content");

	if(cid.val()==0){
		alert("请选择影片分类...");
		cid.focus();
		return false;
	}
	if(name.val().length==0){
		alert("请填写影片中文名...");
		name.focus();
		return false;
	}
	if(ename.val().length==0){
		alert("请填写影片英文名...");
		ename.focus();
		return false;
	}
	if(episode.val().length==0){
		alert("请填写集数...");
		episode.focus();
		return false;
	}
	if(actor.val().length==0){
		alert("请填写主演...");
		actor.focus();
		return false;
	}
	if(director.val().length==0){
		alert("请填写导演...");
		director.focus();
		return false;
	}
	if(updateweek.val().length==0){
		alert("请填写更新星期...");
		updateweek.focus();
		return false;
	}
	if(playdate.val().length==0){
		alert("请填写开播时间...");
		playdate.focus();
		return false;
	}
	if(content.val().length==0){
		alert("请填写影片内容...");
		content.focus();
		return false;
	}
	return true;
}




