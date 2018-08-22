package com.yauritux.service.validator;

import java.math.BigDecimal;

import com.yauritux.exception.CardException;
import com.yauritux.model.constant.CardOperationType;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface CardValidator {

	CardException validateOwner(String owner);
	CardException validateSerialNumber(final String serialNo);
	CardException validateAmount(final BigDecimal amount, final CardOperationType cardOperationType);
}
