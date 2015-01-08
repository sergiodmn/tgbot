-- Copyright 2015 The tgbot Authors. All rights reserved.
-- Use of this source code is governed by a BSD-style
-- license that can be found in the LICENSE file.

started = 0

function sanitize_str(str)
	return string.gsub(str, " ", "_")
end

function get_title(from, to)
	if (to.type == "user") then
		return from.print_name
	elseif (to.type == "chat") then
		return to.print_name
	elseif (to.type == "encr_chat") then
		return from.print_name
	else
		return ""
	end
end

function on_msg_receive (msg)
	if started == 0 then
		return
	end
	print("[MSG] "..
		sanitize_str(get_title(msg.from, msg.to)).." "..
		sanitize_str(msg.from.print_name).." "..
		msg.text)
end

function on_binlog_replay_end ()
	started = 1
end

-- Fix error "*** lua: attempt to call a nil value"
function on_our_id (id) end
function on_user_update (user, what) end
function on_chat_update (chat, what) end
function on_secret_chat_update (schat, what) end
function on_get_difference_end () end
