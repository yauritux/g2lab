package com.yauritux.service.command;

import com.yauritux.model.entity.Card;
import com.yauritux.model.entity.Station;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface CardTransactionCommandService {

	Card entryBus();
	Card entryBarrier(Station station);
	Card exitBarrier(Station station);
}
