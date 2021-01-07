
              <form class="form-horizontal adminex-form" id="documentsub-form">
                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>名称</label>
                  <div class="col-sm-10">
                    <p>{{.doc.Name}}</p>
                  </div>
                </div>
              <div class="form-group">
                <label class="col-sm-2 col-sm-2 control-label"><span>*</span>内容</label>
                <div class="col-sm-10">
                  <p>{{.doc.Content}}</p>
                </div>
              </div>
              <div class="form-group">
                <label class="col-sm-2 col-sm-2 control-label"><span>*</span>要求</label>
                <div class="col-sm-10">
                  <p>{{.doc.Need}}</p>
                </div>
              </div>

                <div class="form-group">
                  <label class="col-sm-2 col-sm-2 control-label"><span>*</span>名称</label>
                  <div class="col-sm-10">
                    <input type="text" name="feedback" value="{{.missionmy.Feedback}}" class="form-control" placeholder="请填写信息">
                  </div>
                </div>
                <div class="form-group">
                <label class="col-sm-2 col-sm-2 control-label"><span>*</span>附件</label>
                <div class="col-sm-10">
                  {{range $k,$v:=.files}}

                  {{if eq $v.Types 1}}
                  <p><a target="_blank" href="{{$v.Path}}"><img width="225" src="{{$v.Path}}"/></a> <span data-id="{{$v.Id}}" class="delete-file fa-2x"> <i style="color: red" class="fa fa-trash-o"></i> </span></p>
                  {{else}}
                  <p><a target="_blank" href="{{$v.Path}}">{{$v.Oldname}}</a><span data-id="{{$v.Id}}" class="delete-file fa-2x"> <i style="color: red" class="fa fa-trash-o"></i> </span></p>
                  {{end}}
                  {{end}}


                  <p><input type="file" name="attachment" multiple="multiple">
                    <span> - </span><br/></p>
                  <p><input type="file" name="attachment" multiple="multiple">
                    <span> - </span><br/></p>
                  <p><input type="file" name="attachment" multiple="multiple">
                    <span> - </span><br/></p>
                </div>
                <div class="form-group">
                  <label class="col-lg-2 col-sm-2 control-label"></label>
                  <div class="col-lg-10">
                    <input type="hidden" name="id" value="270">
                    <button type="submit" class="btn btn-primary">提 交</button>
                  </div>
                </div>
              </form>
